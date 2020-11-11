# Use python container image
FROM python:3.7.2

# Set working directory to /app
WORKDIR /app

# Copy contents of local /app to container's /app
ADD . /app

# Install dependencies within container
RUN pip install -r requirements.txt
RUN pip install requests
RUN pip install --upgrade google-api-python-client google-auth-httplib2 google-auth-oauthlib

# Run command to start uWSGI
CMD ["uwsgi", "app.ini"]