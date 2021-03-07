FROM golang:1.16

#==================
# Environment
#==================
ENV CHROMEDRIVER_VERSION=87.0.4280.20
ENV DISPLAY=:99

#==================
# Apt dependencies
#==================
RUN apt-get update && \
	apt-get install nano git make unzip python-pip xvfb libgbm1 default-jre -y
# python dependencies
RUN pip install selenium chromedriver pyvirtualdisplay

#==================
# Chrome
#==================
RUN apt-get install gconf-service libasound2 libgconf-2-4 libgtk-3-0 libnspr4 \
			libxtst6 fonts-liberation libnss3 lsb-release xdg-utils libxss1 \
			libappindicator1 libindicator7 -y
RUN wget http://dl.google.com/linux/chrome/deb/pool/main/g/google-chrome-stable/google-chrome-stable_87.0.4280.141-1_amd64.deb
RUN dpkg -i google-chrome*.deb

#==================
# ChromeDriver
#==================
RUN wget https://chromedriver.storage.googleapis.com/${CHROMEDRIVER_VERSION}/chromedriver_linux64.zip
RUN unzip chromedriver_linux64.zip
RUN chmod +x chromedriver
RUN mv -f chromedriver /opt/chromedriver
RUN ln -s /opt/chromedriver /usr/bin/chromedriver

#==================
# Virtual Display
#==================
RUN Xvfb :99 &

#==================
# Cleaning up
#==================
RUN rm google-chrome*.deb
RUN rm chromedriver_linux*.zip

# ENVs necessary for selenium service
RUN mkdir /app
COPY ./upwork /app/
WORKDIR /app
ENV SELENIUM_DRIVER=/opt/chromedriver
ENV SELENIUM_JAR=/opt/selenium-server.jar

#==================
# Mocked HTML pages
#==================
COPY ./html_pages/ /html_pages/

#==================
# Selenium Server
#==================
ADD https://selenium-release.storage.googleapis.com/3.141/selenium-server-standalone-3.141.59.jar /opt/selenium-server.jar


# Start
CMD ["/app/upwork"]