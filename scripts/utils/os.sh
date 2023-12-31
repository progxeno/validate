#!/bin/bash
#
# Library for operating system actions.

source "$(dirname "${BASH_SOURCE[0]}")/util.sh"

########################
# Get the name of os.
# Arguments:
#   None
# Returns:
#   Name
#########################
os_get_name() {
  local retval
  retval="$(uname | tr '[:upper:]' '[:lower:]')"

  echo "${retval}"
}

########################
# Check if an user exists in the system.
# Arguments:
#   $1 - user
# Returns:
#   Boolean
#########################
os_exists_user() {
  local user="${1:?user is missing}"
  id "$user" >/dev/null 2>&1
}

########################
# Check if a group exists in the system.
# Arguments:
#   $1 - group
# Returns:
#   Boolean
#########################
os_group_exists() {
  local group="${1:?group is missing}"
  getent group "$group" >/dev/null 2>&1
}

########################
# Create a group in the system if it does not exist already.
# Arguments:
#   $1 - group
# Returns:
#   None
#########################
os_create_group() {
  local group="${1:?group is missing}"

  if ! os_group_exists "$group"; then
    groupadd "$group" >/dev/null 2>&1
  fi
}

########################
# Create an user in the system if it does not exist already.
# Arguments:
#   $1 - user
#   $2 - group (optional)
# Returns:
#   None
#########################
os_create_user() {
  local user="${1:?user is missing}"
  local group="${2:-}"

  if ! os_exists_user "${user}"; then
    useradd "${user}" >/dev/null 2>&1

    if util_is_string "${group}"; then
      os_create_group "${group}"
      usermod -a -G "${group}" "${user}" >/dev/null 2>&1
    fi
  fi
}

########################
# Check if the script is currently running as root.
# Arguments:
#   $1 - user
#   $2 - group
# Returns:
#   Boolean
#########################
os_am_i_root() {
  if [[ "$(id -u)" -ne 0 ]]; then
    return 1
  fi

  return 0
}

########################
# Get total memory available.
# Arguments:
#   None
# Returns:
#   Memory in bytes
#########################
os_get_total_memory() {
  echo $(($(grep MemTotal /proc/meminfo | awk '{print $2}') / 1024))
}

#########################
# Redirects output to /dev/null if debug mode is disabled.
# Globals:
#   DEBUG - (default: false)
# Arguments:
#   $@ - Command to execute
# Returns:
#   None
#########################
os_debug_execute() {
  if ${DEBUG:-false}; then
    "$@"
  else
    "$@" >/dev/null 2>&1
  fi
}

########################
# Retries a command a given number of times.
# Arguments:
#   $1 - cmd as a string
#   $2 - max retries (default: 12)
#   $3 - sleep between retries in seconds (default: 5)
# Returns:
#   Boolean
#########################
os_retry_while() {
  local -r cmd="${1:?cmd is missing}"
  local -r retries="${2:-12}"
  local -r sleep_time="${3:-5}"

  local retval=1

  read -r -a commands <<<"$cmd"
  for ((i = 1; i <= retries; i += 1)); do
    "${commands[@]}" && retval=0 && break
    sleep "$sleep_time"
  done

  return "${retval}"
}
