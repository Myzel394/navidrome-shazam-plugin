# Fetch lyrics for a random song via the Navidrome SubSonic API
[group('test')]
fetch-lyrics:
    #!/usr/bin/env bash
    set -euo pipefail
    auth='u={{username}}&p={{password}}&v=1.16.0&c=test&f=json'
    song_id=$(
        curl -s "http://localhost:4533/rest/getRandomSongs?$auth" \
        | jq -r '.["subsonic-response"].randomSongs.song[0].id'
    )
    echo "$song_id"
    lyrics=$(
        curl -s "http://localhost:4533/rest/getLyricsBySongId?id=$song_id&$auth" \
        
    )
    echo "$lyrics"
