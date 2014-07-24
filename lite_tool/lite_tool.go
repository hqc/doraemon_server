package lite_tool

//from djb2
//unsigned long
//    hash(unsigned char *str)
//    {
//        unsigned long hash = 5381;
//        int c;

//        while (c = *str++)
//            hash = ((hash << 5) + hash) + c; /* hash * 33 + c */

//        return hash;
//    }
func Hash(s string) uint64 {
	h := uint64(5381)
	//var c int;
	for c := range s {
		h = ((h << 5) + h) + uint64(c)
	}
	return h
}
