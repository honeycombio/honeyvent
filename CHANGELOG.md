# Honeyvent Changelog

## 1.1.3 2022-07-20

### Maintenance

- fix openSSL CVE by re-releasing, which will use a new image to build | [@kentquirk](https://github.com/kentquirk)

## 1.1.2 2022-04-26

### Maintenance

- update circle image to go1.18 (#52) | [@JamieDanielson](https://github.com/JamieDanielson)
  - fixes openSSL CVE
- fix github release commands during publish (#53) | [@MikeGoldsmith](https://github.com/MikeGoldsmith)

## 1.1.1 2022-01-10

### Maintenance

- Update go and libhoney-go (#48) | [@MikeGoldsmith](https://github.com/MikeGoldsmith)
- gh: add re-triage workflow (#47) | [@vreynolds](https://github.com/vreynolds)
- gh: add re-triage workflow (#46) | [@vreynolds](https://github.com/vreynolds)
- Update dependabot to monthly (#45) | [@vreynolds](https://github.com/vreynolds)
- create draft github release on publish (#44) | [@MikeGoldsmith](https://github.com/MikeGoldsmith)

## 1.1.0 2021-11-05

### Enhancements

- add --no-tls-verify flag (#41)

### Fixes

- Bump libhoney-go to v1.15.6 (#42)
- empower apply-labels action to apply labels (#40)
- Bump github.com/honeycombio/libhoney-go from 1.15.4 to 1.15.5 (#39)
- Change maintenance badge to maintained (#37)
- Adds Stalebot (#38)
- Add issue and PR templates (#36)
- Add OSS lifecycle badge (#35)
- Add community health files (#34)
