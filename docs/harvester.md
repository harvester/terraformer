### Use with Harvester

Example:

```
 terraformer import harvester -r image,virtualmachine
 terraformer import harvester -r image -f image=default/image-6skr8:harvester-public/image-9lspw
 NAMESPACE=harvester-public terraformer import harvester -r image
```

List of supported Harvester resources:
*   `clusternetwork`
    * `harvester_clusternetwork`
*   `network`
    * `harvester_network`
*   `ssh_key`
    * `harvester_ssh_key`
*   `image`
    * `harvester_image`
*   `virtualmachine`
    * `harvester_virtualmachine`
*   `volume`
    * `harvester_volume`