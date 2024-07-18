# Changelog

## [0.1.7](https://github.com/nobbs/mapr-ticket-parser/compare/v0.1.6...v0.1.7) (2024-07-18)


### Continuous Integration

* add trivy scan to pr workflow ([#47](https://github.com/nobbs/mapr-ticket-parser/issues/47)) ([c5b7c9b](https://github.com/nobbs/mapr-ticket-parser/commit/c5b7c9b22fd9803517a7e14ce9dc56670437ec6a))
* **trivy:** fix token permissions ([e7df43e](https://github.com/nobbs/mapr-ticket-parser/commit/e7df43eab74f3db56c8ee5959cbbf9888c5da3f3))


### Bug Fixes

* **deps:** update k8s.io/utils digest to 18e509b ([#44](https://github.com/nobbs/mapr-ticket-parser/issues/44)) ([e2ca05a](https://github.com/nobbs/mapr-ticket-parser/commit/e2ca05a54e116822dd30ac81bbe10a73f9d0d455))
* **deps:** update module k8s.io/apimachinery to v0.30.3 ([#45](https://github.com/nobbs/mapr-ticket-parser/issues/45)) ([a88c272](https://github.com/nobbs/mapr-ticket-parser/commit/a88c27296ab0fbef160ef0fa59232d3dd107f9d2))
* **python:** update dependencies in `hack/` ([9a880af](https://github.com/nobbs/mapr-ticket-parser/commit/9a880af7b24a5720dedcd2afd61fe4539b347522))

## [0.1.6](https://github.com/nobbs/mapr-ticket-parser/compare/v0.1.5...v0.1.6) (2024-07-06)


### Miscellaneous Chores

* **deps:** update golangci/golangci-lint-action action to v6 ([#38](https://github.com/nobbs/mapr-ticket-parser/issues/38)) ([aa8eb5e](https://github.com/nobbs/mapr-ticket-parser/commit/aa8eb5e772739f9bbc473931ccb2e73c7430110f))
* regenerate from protobuf ([e1d017d](https://github.com/nobbs/mapr-ticket-parser/commit/e1d017d7ba4c72261682855f26843f992c553588))
* switch from asdf to mise, update go to 1.22.5 ([#42](https://github.com/nobbs/mapr-ticket-parser/issues/42)) ([6a9b66e](https://github.com/nobbs/mapr-ticket-parser/commit/6a9b66ef06476dba546b889c6f43c487d7162d50))


### Continuous Integration

* **pre-commit:** update pre-commit hooks and golangci-lint version ([94c9662](https://github.com/nobbs/mapr-ticket-parser/commit/94c96629a596981431b7fccbb585baa3058c965a))


### Bug Fixes

* **deps:** update dependency protobuf to v5.27.2 ([#41](https://github.com/nobbs/mapr-ticket-parser/issues/41)) ([de17fb7](https://github.com/nobbs/mapr-ticket-parser/commit/de17fb77e5d698db759d6536d4a9c57d1885de91))
* **deps:** update k8s.io/utils digest to fe8a2dd ([#40](https://github.com/nobbs/mapr-ticket-parser/issues/40)) ([b83f666](https://github.com/nobbs/mapr-ticket-parser/commit/b83f6666341642fe6945eb6452d759435d40d0e6))
* **deps:** update module google.golang.org/protobuf to v1.34.2 ([#36](https://github.com/nobbs/mapr-ticket-parser/issues/36)) ([d370f14](https://github.com/nobbs/mapr-ticket-parser/commit/d370f14394a8832f6f0d07bb3d0f3ceb48d0441f))
* **deps:** update module k8s.io/apimachinery to v0.30.2 ([#33](https://github.com/nobbs/mapr-ticket-parser/issues/33)) ([4bb5a8c](https://github.com/nobbs/mapr-ticket-parser/commit/4bb5a8c6ee90a4a61cff1ce11089b0314778881a))

## [0.1.5](https://github.com/nobbs/mapr-ticket-parser/compare/v0.1.4...v0.1.5) (2024-04-17)


### Miscellaneous Chores

* **deps:** update dependency golang to v1.22.2 ([#31](https://github.com/nobbs/mapr-ticket-parser/issues/31)) ([775bdb4](https://github.com/nobbs/mapr-ticket-parser/commit/775bdb4c433990df96736e4c6d924a507fae0d4c))
* **deps:** update dependency protobuf to v5 ([#27](https://github.com/nobbs/mapr-ticket-parser/issues/27)) ([8320c8e](https://github.com/nobbs/mapr-ticket-parser/commit/8320c8e02f5a5dfcdb3e11780446327678da1515))


### Bug Fixes

* **deps:** update dependency protobuf to v5.26.1 ([#30](https://github.com/nobbs/mapr-ticket-parser/issues/30)) ([bad05e2](https://github.com/nobbs/mapr-ticket-parser/commit/bad05e2a1b04d98a1284ac4272a52fd5d1c77162))
* **deps:** update k8s.io/utils digest to 4693a02 ([#26](https://github.com/nobbs/mapr-ticket-parser/issues/26)) ([3bc1dde](https://github.com/nobbs/mapr-ticket-parser/commit/3bc1dde8ad248f458c358429be3ac2c51e5815ba))
* **deps:** update module k8s.io/apimachinery to v0.29.3 ([#29](https://github.com/nobbs/mapr-ticket-parser/issues/29)) ([b6dfe2c](https://github.com/nobbs/mapr-ticket-parser/commit/b6dfe2c3bedf74ee4160c4bd813200b4c02df385))
* **deps:** update module k8s.io/apimachinery to v0.29.4 ([#32](https://github.com/nobbs/mapr-ticket-parser/issues/32)) ([1502880](https://github.com/nobbs/mapr-ticket-parser/commit/1502880407fcc541ffbdcfd26c1d755bbfcf855c))

## [0.1.4](https://github.com/nobbs/mapr-ticket-parser/compare/v0.1.3...v0.1.4) (2024-03-05)


### Miscellaneous Chores

* **deps:** refresh sum-file ([9214f00](https://github.com/nobbs/mapr-ticket-parser/commit/9214f00900119993b6cc9b9d942d597dded30156))
* **deps:** update codecov/codecov-action action to v4 ([#16](https://github.com/nobbs/mapr-ticket-parser/issues/16)) ([592dd36](https://github.com/nobbs/mapr-ticket-parser/commit/592dd367e6379ea56d3126a39dfd32aa60137193))
* **deps:** update dependency golang to v1.21.7 ([#17](https://github.com/nobbs/mapr-ticket-parser/issues/17)) ([f1bf4bb](https://github.com/nobbs/mapr-ticket-parser/commit/f1bf4bbf79e8a1b8d9f8d0becdda461a0c4db460))
* **deps:** update dependency golang to v1.22.0 ([#22](https://github.com/nobbs/mapr-ticket-parser/issues/22)) ([b4de28d](https://github.com/nobbs/mapr-ticket-parser/commit/b4de28db77b6395e5a04d48a75851fed45b86065))
* **deps:** update dependency golang to v1.22.1 ([#24](https://github.com/nobbs/mapr-ticket-parser/issues/24)) ([c6daca1](https://github.com/nobbs/mapr-ticket-parser/commit/c6daca173fe91667499030cbc50c657cac231e9a))
* **deps:** update golangci/golangci-lint-action action to v4 ([#19](https://github.com/nobbs/mapr-ticket-parser/issues/19)) ([242d6fb](https://github.com/nobbs/mapr-ticket-parser/commit/242d6fb620b517729d218da520647096020e592b))
* typo in SPDX license header ([31c5b42](https://github.com/nobbs/mapr-ticket-parser/commit/31c5b4231c11ad9ef541d79172fa84730516b4a8))


### Continuous Integration

* update precommit ([65f4066](https://github.com/nobbs/mapr-ticket-parser/commit/65f4066d719bc9abf9efb0734a198a3ca372ff8f))


### Bug Fixes

* **deps:** update dependency protobuf to v4.25.3 ([#21](https://github.com/nobbs/mapr-ticket-parser/issues/21)) ([610d925](https://github.com/nobbs/mapr-ticket-parser/commit/610d925b1ca5d65c4212b1858c4a5c2179574828))
* **deps:** update module github.com/stretchr/testify to v1.9.0 ([#23](https://github.com/nobbs/mapr-ticket-parser/issues/23)) ([01b281b](https://github.com/nobbs/mapr-ticket-parser/commit/01b281b0f3f97199e39ca53483303c531e5254c6))
* **deps:** update module google.golang.org/protobuf to v1.33.0 ([#25](https://github.com/nobbs/mapr-ticket-parser/issues/25)) ([120836e](https://github.com/nobbs/mapr-ticket-parser/commit/120836e5c14bcf60b70a4adc2dff8de7b4cc2ed7))
* **deps:** update module k8s.io/apimachinery to v0.29.2 ([#20](https://github.com/nobbs/mapr-ticket-parser/issues/20)) ([fc4a4ee](https://github.com/nobbs/mapr-ticket-parser/commit/fc4a4eec5e937bfc19c11bdadb4d72b40f34271c))

## [0.1.3](https://github.com/nobbs/mapr-ticket-parser/compare/v0.1.2...v0.1.3) (2024-01-27)


### Miscellaneous Chores

* add license headers to all files ([61ad551](https://github.com/nobbs/mapr-ticket-parser/commit/61ad551ba423527de066eb8baaa0d979f1632401))


### Continuous Integration

* use gotestsum for running tests ([a35059f](https://github.com/nobbs/mapr-ticket-parser/commit/a35059f80852dc17d6120e99b6e0c6879b6e99c8))


### Documentation

* add examples to go docs ([0131d09](https://github.com/nobbs/mapr-ticket-parser/commit/0131d099b2c6fa6af346063fde271e9e729a5d17))
* update README.md ([148f8f5](https://github.com/nobbs/mapr-ticket-parser/commit/148f8f57272badf33f096399134274529d618399))


### Bug Fixes

* add Equal function to MaprTicket type ([3bf2f4a](https://github.com/nobbs/mapr-ticket-parser/commit/3bf2f4a34f6bea4274cdd138bae0abe3c203717d))


### Code Refactoring

* improve testing capabilities for aes package ([a35059f](https://github.com/nobbs/mapr-ticket-parser/commit/a35059f80852dc17d6120e99b6e0c6879b6e99c8))
* simplify pretty printing ([2610b06](https://github.com/nobbs/mapr-ticket-parser/commit/2610b06070144531cedb3b071dfb6f384049de05))
* use proto.Equal in aes_test ([a35059f](https://github.com/nobbs/mapr-ticket-parser/commit/a35059f80852dc17d6120e99b6e0c6879b6e99c8))


### Tests

* add more test for `parse` package ([#12](https://github.com/nobbs/mapr-ticket-parser/issues/12)) ([a35059f](https://github.com/nobbs/mapr-ticket-parser/commit/a35059f80852dc17d6120e99b6e0c6879b6e99c8))
* pin timezone for CI in examples ([ca920c4](https://github.com/nobbs/mapr-ticket-parser/commit/ca920c4031ed897edf40073813f86e9cddf7d626))

## [0.1.2](https://github.com/nobbs/mapr-ticket-parser/compare/v0.1.1...v0.1.2) (2024-01-25)


### Features

* prettify JSON output ([#11](https://github.com/nobbs/mapr-ticket-parser/issues/11)) ([8400884](https://github.com/nobbs/mapr-ticket-parser/commit/8400884758d9f123b75717313bf1f2a007bd435a))


### Bug Fixes

* **deps:** update dependency protobuf to v4.25.2 ([#8](https://github.com/nobbs/mapr-ticket-parser/issues/8)) ([6fc990e](https://github.com/nobbs/mapr-ticket-parser/commit/6fc990ed54b90ff85ea7e61622fcaae919f6e358))
* **deps:** update k8s.io/utils digest to e7106e6 ([#7](https://github.com/nobbs/mapr-ticket-parser/issues/7)) ([a795023](https://github.com/nobbs/mapr-ticket-parser/commit/a795023c36df490e16ec9f8c8ddcec2cc7363ae3))

## [0.1.1](https://github.com/nobbs/mapr-ticket-parser/compare/v0.1.0...v0.1.1) (2024-01-01)


### Miscellaneous Chores

* **deps:** clean up go.mod ([3ee8795](https://github.com/nobbs/mapr-ticket-parser/commit/3ee8795ab7cc6712997a8240b7d522ffe5f690f3))


### Continuous Integration

* enable codecov reporting ([27669e7](https://github.com/nobbs/mapr-ticket-parser/commit/27669e7885049cdcf67f5a9537820d101d1d3851))
* **release:** properly update module version in readme ([73a2542](https://github.com/nobbs/mapr-ticket-parser/commit/73a25421f5c8ded3df87e22810a641097cc9fe17))
* **renovate:** add renovate.json ([#3](https://github.com/nobbs/mapr-ticket-parser/issues/3)) ([8c121f0](https://github.com/nobbs/mapr-ticket-parser/commit/8c121f01df35309745f5e803fefea85d61bba66e))
* **tests:** enable code coverage reporting ([27669e7](https://github.com/nobbs/mapr-ticket-parser/commit/27669e7885049cdcf67f5a9537820d101d1d3851))


### Features

* add function to create new empty ticket ([#6](https://github.com/nobbs/mapr-ticket-parser/issues/6)) ([3ee8795](https://github.com/nobbs/mapr-ticket-parser/commit/3ee8795ab7cc6712997a8240b7d522ffe5f690f3))

## 0.1.0 (2023-12-28)


### Miscellaneous Chores

* add initial files ([de46d17](https://github.com/nobbs/mapr-ticket-parser/commit/de46d177b6b81a3f2b6d9b20b121bd2a33245584))


### Continuous Integration

* add golangci-lint ([ecefa99](https://github.com/nobbs/mapr-ticket-parser/commit/ecefa99e6b12f967bdd3805b301f26f6a07dfd27))
* add release please workflow ([e7cb860](https://github.com/nobbs/mapr-ticket-parser/commit/e7cb860dc9e1815b4f34cc59fdefe82ffc9578f2))
* **release:** set initial version to v0.1.0 ([af113fd](https://github.com/nobbs/mapr-ticket-parser/commit/af113fd2be13d5eb88d79140825a21ca18b5667f))
* run tests on push and PR ([b168823](https://github.com/nobbs/mapr-ticket-parser/commit/b168823fe8f474170c5e73654e9d11180b5c5564))


### Documentation

* add README and example ([df53b63](https://github.com/nobbs/mapr-ticket-parser/commit/df53b63410e87586273a4b6a7aa42e7ff9fb4c4d))


### Code Refactoring

* move around, clean up, make it work ([e2c1850](https://github.com/nobbs/mapr-ticket-parser/commit/e2c18505a4db609f98d5362d2d962c51b6a8c452))
