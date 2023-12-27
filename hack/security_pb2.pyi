from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CredentialsMsg(_message.Message):
    __slots__ = ("uid", "gids", "isRoot", "userName", "clientIp", "isPrivilegedProcess", "logInSeperateAuditPath", "tenantUid", "tenantGids", "encryptedImpersonationMsg", "fromMastGateway", "capabilities")
    UID_FIELD_NUMBER: _ClassVar[int]
    GIDS_FIELD_NUMBER: _ClassVar[int]
    ISROOT_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    CLIENTIP_FIELD_NUMBER: _ClassVar[int]
    ISPRIVILEGEDPROCESS_FIELD_NUMBER: _ClassVar[int]
    LOGINSEPERATEAUDITPATH_FIELD_NUMBER: _ClassVar[int]
    TENANTUID_FIELD_NUMBER: _ClassVar[int]
    TENANTGIDS_FIELD_NUMBER: _ClassVar[int]
    ENCRYPTEDIMPERSONATIONMSG_FIELD_NUMBER: _ClassVar[int]
    FROMMASTGATEWAY_FIELD_NUMBER: _ClassVar[int]
    CAPABILITIES_FIELD_NUMBER: _ClassVar[int]
    uid: int
    gids: _containers.RepeatedScalarFieldContainer[int]
    isRoot: bool
    userName: str
    clientIp: int
    isPrivilegedProcess: bool
    logInSeperateAuditPath: bool
    tenantUid: int
    tenantGids: _containers.RepeatedScalarFieldContainer[int]
    encryptedImpersonationMsg: bytes
    fromMastGateway: bool
    capabilities: Capabilities
    def __init__(self, uid: _Optional[int] = ..., gids: _Optional[_Iterable[int]] = ..., isRoot: bool = ..., userName: _Optional[str] = ..., clientIp: _Optional[int] = ..., isPrivilegedProcess: bool = ..., logInSeperateAuditPath: bool = ..., tenantUid: _Optional[int] = ..., tenantGids: _Optional[_Iterable[int]] = ..., encryptedImpersonationMsg: _Optional[bytes] = ..., fromMastGateway: bool = ..., capabilities: _Optional[_Union[Capabilities, _Mapping]] = ...) -> None: ...

class Capabilities(_message.Message):
    __slots__ = ("clusterOpsMask",)
    CLUSTEROPSMASK_FIELD_NUMBER: _ClassVar[int]
    clusterOpsMask: int
    def __init__(self, clusterOpsMask: _Optional[int] = ...) -> None: ...

class Key(_message.Message):
    __slots__ = ("key",)
    KEY_FIELD_NUMBER: _ClassVar[int]
    key: bytes
    def __init__(self, key: _Optional[bytes] = ...) -> None: ...

class TicketAndKey(_message.Message):
    __slots__ = ("encryptedTicket", "userKey", "userCreds", "expiryTime", "creationTimeSec", "maxRenewalDurationSec", "canUserImpersonate", "ips", "impersonatedUids", "impersonatedGids", "isTenant", "isExternal", "lastRenewalTime", "canUserGenerateTicket", "isRemoteTempTicket")
    ENCRYPTEDTICKET_FIELD_NUMBER: _ClassVar[int]
    USERKEY_FIELD_NUMBER: _ClassVar[int]
    USERCREDS_FIELD_NUMBER: _ClassVar[int]
    EXPIRYTIME_FIELD_NUMBER: _ClassVar[int]
    CREATIONTIMESEC_FIELD_NUMBER: _ClassVar[int]
    MAXRENEWALDURATIONSEC_FIELD_NUMBER: _ClassVar[int]
    CANUSERIMPERSONATE_FIELD_NUMBER: _ClassVar[int]
    IPS_FIELD_NUMBER: _ClassVar[int]
    IMPERSONATEDUIDS_FIELD_NUMBER: _ClassVar[int]
    IMPERSONATEDGIDS_FIELD_NUMBER: _ClassVar[int]
    ISTENANT_FIELD_NUMBER: _ClassVar[int]
    ISEXTERNAL_FIELD_NUMBER: _ClassVar[int]
    LASTRENEWALTIME_FIELD_NUMBER: _ClassVar[int]
    CANUSERGENERATETICKET_FIELD_NUMBER: _ClassVar[int]
    ISREMOTETEMPTICKET_FIELD_NUMBER: _ClassVar[int]
    encryptedTicket: bytes
    userKey: Key
    userCreds: CredentialsMsg
    expiryTime: int
    creationTimeSec: int
    maxRenewalDurationSec: int
    canUserImpersonate: bool
    ips: _containers.RepeatedScalarFieldContainer[int]
    impersonatedUids: _containers.RepeatedScalarFieldContainer[int]
    impersonatedGids: _containers.RepeatedScalarFieldContainer[int]
    isTenant: bool
    isExternal: bool
    lastRenewalTime: int
    canUserGenerateTicket: bool
    isRemoteTempTicket: bool
    def __init__(self, encryptedTicket: _Optional[bytes] = ..., userKey: _Optional[_Union[Key, _Mapping]] = ..., userCreds: _Optional[_Union[CredentialsMsg, _Mapping]] = ..., expiryTime: _Optional[int] = ..., creationTimeSec: _Optional[int] = ..., maxRenewalDurationSec: _Optional[int] = ..., canUserImpersonate: bool = ..., ips: _Optional[_Iterable[int]] = ..., impersonatedUids: _Optional[_Iterable[int]] = ..., impersonatedGids: _Optional[_Iterable[int]] = ..., isTenant: bool = ..., isExternal: bool = ..., lastRenewalTime: _Optional[int] = ..., canUserGenerateTicket: bool = ..., isRemoteTempTicket: bool = ...) -> None: ...
