package algorithm

import "ultipa-go-sdk/sdk"

// NewKHopParams create KHop params for Page Rank task
func NewEmbeddingParams() TaskValues {
	return TaskValues{
		"walk_num":              "1",
		"walk_length":           "3",
		"p":                     "1",
		"q":                     "0",
		"context_size":          "100",
		"dimension_word_vector": "100",
		"learning_rate":         "0.025",
		"min_learning_rate":     "0.0001",
		"min_frequency":         "0",
		"sample":                "-1",
		"resolution":            "10",
		"neg_number":            "1",
		"iter_num":              "5",
	}
}

//  * ----------parameters for n2v walk----------
//  * walk_num : number of the walk for each node
//  * walk_length : step number for each walk
//  * p : Return parameter , p越大，往回走的概率越小，反之越高
//  * q : In-out papameter , 控制着游走是向外还是向内，若 [q>1] ，随机游走倾向于访问和当前点接近的顶点(偏向BFS)。若 [q<1] ，倾向于访问远离当前点的顶点(偏向DFS)
//  *
//  * ----------parameters for embedding---------
//  * context_size : window size for the word2vec
//  * dimension_word_vector : dimension of the embedded vector , 10~n100  (e.g : 100)
//  * learning_rate : for SGD in the word2vec, 0.025 recommended
//  * min_learning_rate : learning_rate * 0.0001 recommended;
//  * min_frequency : used to remove words that occur fewer than 'min_frequency' times in the training text , ps : 0 or 1 means no removing and 0 is recommended
//  * sample : -1 default,  ps : can be 0.001 in word2vec, the equation is (sqrt(x / 0.001) + 1) * (0.001 / x)
//  * resolution : resolution * vocab_size is the size of the unigramTable   (10 or 100 recommended)
//  * neg_number : number of words in negative sampling phrase
//  * iter_num : iteration number of the main loop in the word2vec

// StartCCTask starts a Page Rank task
func StartEmbeddingTask(client sdk.Client, params TaskValues) *TaskReply {
	return StartTask(client, TaskEmbedding, params)
}
