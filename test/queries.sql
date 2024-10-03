SELECT 'query: highest temperature ever';
SELECT
    tempMax / 10 AS maxTemp,
    location,
    name,
    date
FROM noaa.ghcnd
WHERE tempMax > 500
ORDER BY
    tempMax DESC,
    date ASC
LIMIT 5;

SELECT 'query: best ski resorts';
SELECT
    resort_name,
    total_snow / 1000 AS total_snow_m,
    resort_location,
    month_year
FROM
    (
        WITH resorts AS
                 (
                     SELECT
                         resort_name,
                         state,
                         (lon, lat) AS resort_location,
                         'US' AS code
                     FROM url('https://gist.githubusercontent.com/gingerwizard/dd022f754fd128fdaf270e58fa052e35/raw/622e03c37460f17ef72907afe554cb1c07f91f23/ski_resort_stats.csv', CSVWithNames)
                 )
        SELECT
            resort_name,
            highest_snow.station_id,
            geoDistance(resort_location.1, resort_location.2, station_location.1, station_location.2) / 1000 AS distance_km,
            highest_snow.total_snow,
            resort_location,
            station_location,
            month_year
        FROM
            (
                SELECT
                    sum(snowfall) AS total_snow,
                    station_id,
                    any(location) AS station_location,
                    month_year,
                    substring(station_id, 1, 2) AS code
                FROM noaa.ghcnd
                WHERE (date > '2017-01-01') AND (code = 'US') AND (elevation > 1800)
                GROUP BY
                    station_id,
                    toYYYYMM(date) AS month_year
                ORDER BY total_snow DESC
                LIMIT 1000
                ) AS highest_snow
                INNER JOIN resorts ON highest_snow.code = resorts.code
        WHERE distance_km < 20
        ORDER BY
            resort_name ASC,
            total_snow DESC
        LIMIT 1 BY
            resort_name,
            station_id
        )
ORDER BY total_snow DESC
LIMIT 5;
