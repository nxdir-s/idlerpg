yq -i '(select(.data[]) | .data[]) style="double"' ./dist/0002-idlerpg.k8s.yaml
