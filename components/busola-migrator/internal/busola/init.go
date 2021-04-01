package busola

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"k8s.io/client-go/rest"

	"github.com/kyma-project/kyma/components/busola-migrator/internal/config"
)

const initTemplate string = `{
  "cluster": {
    "server": "%s",
    "certificate-authority-data": "%s"
  },
  "auth": {
    "issuerUrl": "%s",
    "clientId": "%s",
    "scope": "%s",
    "usePKCE": %t
  },
  "config": {
    "disabledNavigationNodes": "",
    "systemNamespaces": "istio-system knative-eventing knative-serving kube-public kube-system kyma-backup kyma-installer kyma-integration kyma-system natss kube-node-lease kubernetes-dashboard serverless-system"
  },
  "features": {
    "bebEnabled": false
  }
}`

func BuildInitURL(appCfg config.Config, kubeConfig *rest.Config) (string, error) {
	if kubeConfig == nil {
		return "", errors.New("Kubeconfig not found")
	}

	host, err := url.Parse(kubeConfig.Host)
	if err != nil {
		return "", errors.New("host not found")
	}

	initString := fmt.Sprintf(initTemplate,
		host.Hostname(),
		base64.StdEncoding.EncodeToString(kubeConfig.CAData),
		appCfg.OIDCIssuerURL,
		appCfg.OIDCClientID,
		appCfg.OIDCScope,
		appCfg.OIDCUsePKCE,
	)
	fmt.Println(kubeConfig)
	fmt.Println(initString)
	encodedInitString := encodeInitString(initString)
	fmt.Println(encodedInitString)

	initURL, err := url.ParseRequestURI(fmt.Sprintf("%s/?init=%s", appCfg.BusolaURL, encodedInitString))
	if err != nil {
		return "", errors.Wrap(err, "Busola init url couldn't be parsed")
	}

	return initURL.String(), nil
}
