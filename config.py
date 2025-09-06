from os import environ

BASE_URL = environ.get('ERO_BASE_URL', 'http://localhost:5000')
INACTIVE_TIMELIMIT = int(environ.get('ERO_INACTIVE_TIMELIMIT', 60*60*4))
API_SECRET = environ.get('ERO_API_SECRET', '')
LOG_PATH = environ.get('ERO_LOG_PATH', '/var/log/ero-runner')
DATA_PATH = environ.get('ERO_DATA_PATH', '/var/lib/ero-runner')

CPU_PERIOD = int(environ.get('ERO_CPU_PERIOD', 100_000))
CPU_QUOTA = int(environ.get('ERO_CPU_QUOTA', 50_000))
MEM_LIMIT = environ.get('ERO_MEM_LIMIT', '512m')
STORAGE_LIMIT = environ.get('ERO_STORAGE_LIMIT', '5G')
