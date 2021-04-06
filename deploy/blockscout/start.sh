#!/bin/sh
printenv
mix do ecto.drop --force, ecto.create, ecto.migrate
mix phx.server