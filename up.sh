docker run -d \
 --name site-duckgame \
 -p 3825:3825 \
 --network prod \
 --restart unless-stopped \
 site-duckgame