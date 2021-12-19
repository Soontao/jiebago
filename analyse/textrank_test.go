package analyse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sentence = "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"
)

func TestTextRank(t *testing.T) {

	assert := assert.New(t)
	var tr TextRanker
	tr.LoadDictionary("../dict.txt")
	results := tr.TextRank(sentence, 10)
	for _, tw := range results {
		assert.Greater(len(tw.Text()), int(0))
		assert.Greater(tw.Weight(), float64(0))
	}
}
