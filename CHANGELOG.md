# Changelog

All notable changes to *Component Metadata* will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and *Component Metadata* adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

> NOTE: Base version is set to 0.0.0 to prevent build. To
> trigger a build commit with updated CHANGELOG.md and bump-up app version in
> VERSION.ini file.

## [1.0.0] - 2021-12-27

- First release
- Service refactored using Cookiecutter Service Template version 1.0.0
- Uses DAL package to manage component contracts used by designer
- Netlist component parameterizes python netlist from blob and returns SPICE netlist in json

## [1.0.1] - 2022-01-13

- Added People Detection Metadata
- Endpoint to get People Detection Metadata to load as initial state on the UI

## [1.0.2] - 2022-01-18

- Modified People Detection Metadata Component to be a generic App Metadata Component
- appId is used as a path parameter to load initial state of particular application on the UI
