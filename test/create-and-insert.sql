CREATE DATABASE IF NOT EXISTS noaa;

CREATE OR REPLACE TABLE noaa.ghcnd
(
    `station_id` LowCardinality(String),
    `date` Date32,
    `tempAvg` Int32 COMMENT 'Average temperature (tenths of a degrees C)',
    `tempMax` Int32 COMMENT 'Maximum temperature (tenths of degrees C)',
    `tempMin` Int32 COMMENT 'Minimum temperature (tenths of degrees C)',
    `precipitation` UInt32 COMMENT 'Precipitation (tenths of mm)',
    `snowfall` UInt32 COMMENT 'Snowfall (mm)',
    `snowDepth` UInt32 COMMENT 'Snow depth (mm)',
    `percentDailySun` UInt8 COMMENT 'Daily percent of possible sunshine (percent)',
    `averageWindSpeed` UInt32 COMMENT 'Average daily wind speed (tenths of meters per second)',
    `maxWindSpeed` UInt32 COMMENT 'Peak gust wind speed (tenths of meters per second)',
    `weatherType` Enum8('Normal' = 0, 'Fog' = 1, 'Heavy Fog' = 2, 'Thunder' = 3, 'Small Hail' = 4, 'Hail' = 5, 'Glaze' = 6, 'Dust/Ash' = 7, 'Smoke/Haze' = 8, 'Blowing/Drifting Snow' = 9, 'Tornado' = 10, 'High Winds' = 11, 'Blowing Spray' = 12, 'Mist' = 13, 'Drizzle' = 14, 'Freezing Drizzle' = 15, 'Rain' = 16, 'Freezing Rain' = 17, 'Snow' = 18, 'Unknown Precipitation' = 19, 'Ground Fog' = 21, 'Freezing Fog' = 22),
    `location` Point,
    `elevation` Float32,
    `name` LowCardinality(String)
) ENGINE = MergeTree() ORDER BY (station_id, date);

INSERT INTO noaa.ghcnd SELECT *
FROM s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/noaa/noaa_enriched.parquet')
