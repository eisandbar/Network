FROM redis
COPY redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "--save 60 1 --loglevel warning /usr/local/etc/redis/redis.conf" ]