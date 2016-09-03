# honeyvent
CLI for sending individual events in to Honeycomb

Use - call with a collection of names and values to send an event from the
command line:

honeyvent -k <writekey> -d <dataset> -n field -v val -n field -v val ...

The tool will detect floats and ints and send them as numbers; everything else
turns in to strings.  Quote any values that have spaces.
