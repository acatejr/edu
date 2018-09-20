#!/bin/bash
docker run -v $PWD:/home/jovyan/work -d -p 8888:8888 jupyter/datascience-notebook
