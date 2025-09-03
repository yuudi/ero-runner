from os import environ

BASE_URL = environ.get('BASE_URL', 'http://localhost:5000')
INACTIVE_TIMELIMIT = int(environ.get('INACTIVE_TIMELIMIT', 60*60*4))
API_SECRET = environ.get('API_SECRET', '')
LOG_PATH = environ.get('LOG_PATH', '/var/log/ero-runner')
DATA_PATH = environ.get('DATA_PATH', '/var/lib/ero-runner')
