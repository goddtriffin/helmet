# Helmet

HTTP security middleware for [Go(lang)](https://golang.org/) inspired by [HelmetJS](https://helmetjs.github.io/).

| Module                                                                                                               | Default Value(s) |
| -------------------------------------------------------------------------------------------------------------------- | ---------------- |
| [Content-Security-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP)                                     |                  |
| [X-DNS-Prefetch-Control](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-DNS-Prefetch-Control)           | `off`            |
| [Expect-CT](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect-CT)                                     |                  |
| [Feature-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Feature-Policy)                           |                  |
| [X-Frame-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options)                         | `SAMEORIGIN`     |
| [X-Permitted-Cross-Domain-Policies](https://owasp.org/www-project-secure-headers/#x-permitted-cross-domain-policies) |                  |
