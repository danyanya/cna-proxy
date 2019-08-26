# CNA Proxy 

Simple HTTP proxy for iOS Captive Network Assistant bypass.

## Logic

Server handle only `/hotspot-*.*` URI - it's common URI for Apple devices to check Internet connectivity.

On first HTTP request from Client IP cna-proxy returns Empty HTML -- this response forced iOS to open CNA splash page.

On second HTTP request from same Client IP cna-proxy returns Special HTML with Success body and redirect to http page -- it forced iOS to open redirect URL in Safari!

## Usage

1. Run on some server IP on 80 port in local network (LAN) with clients
2. Add IP to Captive Portal Walled garden
3. Resolve IP as captive.apple.com on your local DNS server
4. Voila - all checks from clients iOS devices will be handled on cna-proxy
