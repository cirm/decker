#!/usr/bin/env bash
migrate -url postgres://spark:salasala@postgres1.cydec/arco?sslmode=disable --path ./postgres/ up