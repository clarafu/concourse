#### <sub><sup><a name="5770" href="#5770">:link:</a></sup></sub> fix

* `fly login` can now accept arbitrarily long tokens when pasting the token manually into the console. Previously, the limit was OS dependent (with OSX having a relatively small maximum length of 1024 characters). This has been a long-standing issue, but it became most noticable after 6.1.0 which significantly increased the size of tokens. #5770
