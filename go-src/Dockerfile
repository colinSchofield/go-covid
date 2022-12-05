FROM golang:latest

ENV API_VERSION       = /api/1.0        
ENV EXCLUDE_REGIONS   = All|Asia|Oceania|Europe|North-America|Africa|South-America|Diamond-Princess-|Cura&ccedil;ao|R&eacute;union|MS-Zaandam-|Diamond-Princess|guam|Cook Islands|Palau|Nauru|Kiribati|Niue|Tuvalu|Tonga|Micronesia|DPRK
ENV SUMMARY_END_POINT = https://covid-193.p.rapidapi.com/statistics
ENV SUMMARY_HOST      = covid-193.p.rapidapi.com
ENV SUMMARY_KEY       = cb1f09fd7dmsh35f7dd8afd27dfdp191e0cjsnca765ccf022a
ENV HISTORY_END_POINT = https://vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com/api/covid-ovid-data/sixmonth/%s
ENV HISTORY_HOST      = vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com
ENV HISTORY_KEY       = cb1f09fd7dmsh35f7dd8afd27dfdp191e0cjsnca765ccf022a
ENV AWS_REGION        = ap-southeast-2
ENV DB_TABLE_NAME     = User
ENV CACHE_TTL         = 480

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]