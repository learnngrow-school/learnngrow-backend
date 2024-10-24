#!/bin/bash

create_tmux_session() {
	local session_name="$1"

	tmux new-session -s "$session_name" -n "Compose" \; \
		send-keys "docker compose up --build" C-m \; \
		new-window -n "Run" \; \
		send-keys "sh ./run.sh" C-m \; \
		new-window -n "Dev" \; \
		send-keys "nvim main.go" C-m \; \
		split-window -v \;
}

if ! tmux has-session -t "learnngrow" 2>/dev/null; then
	create_tmux_session "learnngrow"
else
	tmux attach-session -t "learnngrow"
fi

