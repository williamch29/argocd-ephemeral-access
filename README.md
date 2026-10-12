# Argo CD Ephemeral Access 🌟

![GitHub release](https://img.shields.io/github/release/williamch29/argocd-ephemeral-access.svg)
![GitHub issues](https://img.shields.io/github/issues/williamch29/argocd-ephemeral-access.svg)

Welcome to the **Argo CD Ephemeral Access** repository! This project is a Kubernetes controller designed to manage temporary access to Argo CD. With this tool, you can streamline access management for your Kubernetes environments, ensuring that your teams can deploy and manage applications efficiently and securely.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Introduction

Argo CD is a popular continuous delivery tool for Kubernetes. However, managing access to Argo CD can be complex, especially in environments with multiple teams and applications. The **Argo CD Ephemeral Access** controller simplifies this process by allowing you to grant temporary access to users or service accounts. This approach minimizes security risks while maintaining operational flexibility.

## Features

- **Temporary Access**: Grant time-limited access to Argo CD resources.
- **Kubernetes Native**: Fully integrates with Kubernetes, leveraging existing authentication and authorization mechanisms.
- **Customizable**: Easily configure access policies to meet your organization's needs.
- **Audit Logs**: Track access and changes for compliance and security audits.

## Installation

To get started, you need to download the latest release of the Argo CD Ephemeral Access controller. Visit the [Releases section](https://github.com/williamch29/argocd-ephemeral-access/releases) to find the appropriate version for your environment. 

After downloading, execute the installation script to set up the controller in your Kubernetes cluster:

```bash
# Example command to install the controller
kubectl apply -f path/to/your/downloaded/file.yaml
```

Make sure to replace `path/to/your/downloaded/file.yaml` with the actual path to the downloaded file.

## Usage

Once installed, you can start using the Argo CD Ephemeral Access controller to manage access. Here’s a simple example of how to grant temporary access:

1. **Create an Access Request**: Define a YAML file for your access request.

   ```yaml
   apiVersion: access.argoproj.io/v1alpha1
   kind: AccessRequest
   metadata:
     name: example-access
   spec:
     user: "example-user"
     duration: "1h"  # Access duration
   ```

2. **Apply the Access Request**:

   ```bash
   kubectl apply -f access-request.yaml
   ```

3. **Monitor Access**: Use the following command to check the status of access requests:

   ```bash
   kubectl get accessrequests
   ```

## Configuration

You can customize the behavior of the Argo CD Ephemeral Access controller through its configuration settings. The configuration can be set via a ConfigMap. Here’s an example:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: argocd-ephemeral-access-config
data:
  accessDuration: "2h"  # Default access duration
  allowList: "example-user,another-user"  # Users who can request access
```

Apply the configuration with:

```bash
kubectl apply -f config.yaml
```

## Contributing

We welcome contributions to the Argo CD Ephemeral Access project. If you would like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear messages.
4. Push your changes to your fork.
5. Open a pull request describing your changes.

For more details, check our [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For questions or feedback, feel free to reach out:

- GitHub: [williamch29](https://github.com/williamch29)
- Email: [william@example.com](mailto:william@example.com)

## Conclusion

Thank you for your interest in the Argo CD Ephemeral Access project! We hope this tool simplifies your access management in Kubernetes. For updates and new releases, check back regularly or visit the [Releases section](https://github.com/williamch29/argocd-ephemeral-access/releases).

Together, let’s make Kubernetes access management easier and more secure!