FROM mcr.microsoft.com/devcontainers/universal:linux

WORKDIR /runner-admin

COPY . .

RUN pip install -r requirements.txt

ENTRYPOINT ["sh", "-c", "gunicorn -w 1 -b 0.0.0.0 main:app"]
