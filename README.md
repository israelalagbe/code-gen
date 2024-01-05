# Code Generator for REST APIs

## Description

This project is an internal tool designed to streamline the process of generating basic REST API code. It automates the creation of common code structures, reducing the time and effort required to set up a new API. Whether you're creating a new microservice or adding to an existing project, this tool can help you get up and running quickly. Simply modify the templates to match your project folder structure and code structure.

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/israelalagbe/code-gen.git
```

Navigate to the project directory:

```bash
cd code-gen
```

## Build

To build the project, use the provided Makefile:

```bash
make build
```

Your binary should be generated in the bin folder

## Usage

Add the binary directory to your path and you can run the command below anywhere you want to use it

```bash
code-gen
```

It will ask you a series of questions about the code you wanted to generate
