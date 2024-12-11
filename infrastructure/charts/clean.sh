yq -i '(select(.data[]) | .data[]) style="double"' ./dist/idlerpg.k8s.yaml
