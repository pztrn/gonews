# NNTP implementation status

This document tracks implementation status for NTTP protocol and extensions. See:

* [RFC3977 - Network News Transfer Protocol (NNTP)](https://tools.ietf.org/html/rfc3977)
* [RFC4643 - Network News Transfer Protocol (NNTP). Extension for Authentication](https://tools.ietf.org/html/rfc4643)
* [RFC6048 - Network News Transfer Protocol (NNTP) Additions to LIST Command](https://tools.ietf.org/html/rfc6048)

After all checkmarks will be checked - implementation can be considered complete.

This document will be updated with new RFCs data in process of development.

RFC3977:

* [ ] Connection handling ([section 5](https://tools.ietf.org/html/rfc3977#section-5)):
  * [x] Concurrent connection handling
  * [x] CAPABILITIES command ([section 5.2](https://tools.ietf.org/html/rfc3977#section-5.2))
  * [ ] MODE command ([section 5.3](https://tools.ietf.org/html/rfc3977#section-5.3))
  * [x] QUIT command ([section 5.4](https://tools.ietf.org/html/rfc3977#section-5.4))
* [ ] Article retrieval and posting, working with groups ([section 6](https://tools.ietf.org/html/rfc3977#section-6)):
  * [ ] Group selection and information retrieval ([section 6.1.1](https://tools.ietf.org/html/rfc3977#section-6.1.1))
  * [ ] Group selection and information retrieval with articles numbers ([section 6.1.2](https://tools.ietf.org/html/rfc3977#section-6.1.2))
  * [ ] LAST command ([section 6.1.3](https://tools.ietf.org/html/rfc3977#section-6.1.3))
  * [ ] NEXT command ([section 6.1.4](https://tools.ietf.org/html/rfc3977#section-6.1.4)
  * [ ] ARTICLE command ([section 6.2.1](https://tools.ietf.org/html/rfc3977#section-6.2.1))
  * [ ] HEAD command ([section 6.2.2](https://tools.ietf.org/html/rfc3977#section-6.2.2))
  * [ ] BODY command ([section 6.2.3](https://tools.ietf.org/html/rfc3977#section-6.2.3))
  * [ ] STAT command ([section 6.2.4](https://tools.ietf.org/html/rfc3977#section-6.2.4))
  * [ ] POST command ([section 6.3.1](https://tools.ietf.org/html/rfc3977#section-6.3.1))
  * [ ] IHAVE command ([section 6.3.2](https://tools.ietf.org/html/rfc3977#section-6.3.2))
* [ ] Informational ([section 7](https://tools.ietf.org/html/rfc3977#section-7)):
  * [ ] DATE command ([section 7.1](https://tools.ietf.org/html/rfc3977#section-7.1))
  * [ ] HELP command ([section 7.2](https://tools.ietf.org/html/rfc3977#section-7.2))
  * [ ] NEWGROUPS command ([section 7.3](https://tools.ietf.org/html/rfc3977#section-7.3))
  * [ ] NEWNEWS command ([section 7.4](https://tools.ietf.org/html/rfc3977#section-7.4))
  * [ ] LIST commands ([section 7.6](https://tools.ietf.org/html/rfc3977#section-7.6))
