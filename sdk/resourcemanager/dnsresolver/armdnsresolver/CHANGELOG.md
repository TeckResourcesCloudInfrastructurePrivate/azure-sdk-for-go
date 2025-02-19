# Release History

## 1.0.0 (2022-09-15)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed
- Struct `ProxyResource` has been removed
- Struct `Resource` has been removed
- Struct `TrackedResource` has been removed

### Features Added

- New field `DNSResolverOutboundEndpoints` in struct `DNSForwardingRulesetPatch`


## 0.4.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.4.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).