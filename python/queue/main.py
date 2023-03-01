from mymodule import count_words_at_url
from redis import Redis
from rq import Queue


q = Queue(connection=Redis())
job = q.enqueue(count_words_at_url, 'http://nvie.com')
