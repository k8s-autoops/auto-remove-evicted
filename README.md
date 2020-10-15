# auto-remove-evicted

## Usage

Create namespace `autoops` and apply yaml resources as described below.

```yaml
# create serviceaccount
apiVersion: v1
kind: ServiceAccount
metadata:
  name: auto-remove-evicted
  namespace: autoops
---
# create clusterrole
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: auto-remove-evicted
rules:
  - apiGroups: [""]
    resources: ["pods"]
    verbs: ["list", "delete"]
---
# create clusterrolebinding
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: auto-remove-evicted
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: auto-remove-evicted
subjects:
  - kind: ServiceAccount
    name: auto-remove-evicted
    namespace: autoops
---
# create cronjob
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: auto-remove-evicted
  namespace: autoops
spec:
  schedule: "*/5 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccount: auto-remove-evicted
          containers:
            - name: auto-remove-evicted
              image: autoops/auto-remove-evicted
          restartPolicy: OnFailure
```

## Credits

Guo Y.K., MIT License
