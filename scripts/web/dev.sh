~/go/bin/air -c .air.toml & \
npx tailwind \
  -i 'web/static/input.css' \
  -o 'web/static/output.css' \
  --watch & \
npx browser-sync start \
  --files 'web/tmpl/**/*.html, web/static/**/*.css' \
  --port 3001 \
  --proxy '127.0.0.1:8080' \
  --middleware 'function(req, res, next) { \
    res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
    return next(); \
  }'

