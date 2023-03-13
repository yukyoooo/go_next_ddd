package enum

type Role int

const (
	Waiting Role = iota
	Working
	ReviewRequest
	InReview
	completed
)

/*
func (r Role) String() string {
	switch r {
	case Waiting:
		return "未着手"
	case Working:
		return "作業中"
	case ReviewRequest:
		return "レビュー依頼中"
	case InReview:
		return "レビュー中"
	case completed:
		return "完了"
	default:
		return "未定義の役割"
	}
}
*/
