# Run Nginx
```
docker run --name nginx -v /Users/golfapipol/sck/GOW/src/gow-frontend/ui.conf:/etc/nginx/conf.d/default.conf:ro -v /Users/golfapipol/sck/GOW/src/gow-frontend/public/:/usr/share/nginx/html:ro --link datecal:gow-backend  -p 8080:80 -d nginx:stable
```