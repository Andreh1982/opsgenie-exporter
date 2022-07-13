# Prometheus Opsgenie Exporter

A simple exporter who gives metrics from incidents total count by status(closed, resolved, opened) and list Postmortems total count by status(closed and resolved). Metrics Type are Gauge.

**TO-DO:**

List Incidents by Teams;
List Incidents by Service;
Alerts Counters by Status;
And more...

## Setup exporter

Create the environment variables like the example bellow:

    ENVIRONMENT=staging
    OPSGENIE_API_URL=https://api.opsgenie.com/v1/
    OPSGENIE_API_KEY="GenieKey ********-****-****-****-*********"

For local use, create a **.env.local** file with those variables in the root folder.

## Run the exporter

Using **go run**:

    go run cmd/opsgenie-exporter/main.go

Using **docker**:

    docker build .
    docker run -it -p 2112:2112 <IMAGE_NAME>

Using **docker-compose**:

    docker-compose up -d

## Prometheus Setup

Add the following job to prometheus scrap configuration:

    - job_name: opsgenie-exporter
        scrape_interval: 15s
        static_configs:
        - targets: 
          - IP-ADDRESS:2112

## Exposed metrics

    opsgenie_incidents_closed
    opsgenie_incidents_resolved
    opsgenie_incidents_opened
    opsgenie_postmortem_incidents_closed
    opsgenie_postmortem_incidents_resolved
    opsgenie_teams_total
