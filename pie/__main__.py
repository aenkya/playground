import os
import sys
__cwd = os.path.abspath(os.path.dirname(__file__))
__modules = os.path.abspath(__cwd)

sys.path.append(os.path.abspath(__modules))

from algorithms import streams

Streams = streams.Streams

def main():

    stream = [
        "host|target|hero:test, hero:view",
        "host|target|hero:test-2, hero:view, hero:all",
        "host|target|hero:test, hero:view, hero:all",
        "host|target|hero:all"
    ]
    inputs = ['hero:all']
    streams  = Streams('input', 'output')
    print(streams.get_stream_tags_from_input(stream, inputs))

if __name__ == "__main__":
    main()