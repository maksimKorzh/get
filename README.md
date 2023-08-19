# get
Retrieving files using HTTP client with a custom DNS resolver<br>

# Purpose
If you have a Linux System with a statically linked BusyBox<br>
running wget command would result in error because of the lack<br>
of libnss_dns.so shared library responsible for DNS resolving.<br>
This utility is a pure static binary workaround.

# Usage
    get                                    # URL via STDIN goes to STDOUT
    get https://github.com                 # URL via  ARG  goes to STDOUT
    get https://github.com > index.html    # URL via  ARG  goes to file 'index.html'
