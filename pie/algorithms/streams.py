
class ProcessingMixin(object):

    @staticmethod
    def get_tags(stream):
        try:
            x = stream.strip().lower().split('|')[2]
            if x is None or len(x) == 0:
                return []
            return [v.strip() for v in x.split(',')]
        except Exception as e:
            print(e, stream)
            return []

    def __init__(self, *args, **kwargs):
        self.args = args
        self.kwargs = kwargs

    def checkIfTagsInLog(self, logs, tags):
        return True if all([self.contains(logs, tag) for tag in tags]) else False
    
    def contains(self, logs, tag):
        return True if tag in logs else False
    
    def removeTags(self, log, tags):
        return [tag for tag in log if tag not in tags]

        

class Streams(ProcessingMixin):
    def __init__(self, *args, **kwargs):
        super(Streams, self).__init__(*args, **kwargs)
        self.store = {}

    def get_stream_tags_from_input(self, streams=[], input_tags=[]):
        for stream in streams:
            log_tags = ProcessingMixin.get_tags(stream)
            if self.checkIfTagsInLog(log_tags, input_tags):
                for tag in log_tags:
                    if tag not in self.store:
                        self.store[tag] = 1
                        continue
                    self.store[tag] += 1
        return self.removeTags(self.store.keys(), input_tags)

    

            

    
    
    
    