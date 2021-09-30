package pkg

type Activity struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Label       string `json:"label"`
	Deadline    string `json:"deadline"`
	Severity    string `json:"severity"`
	Priority    string `json:"priority"`
}

type ActivityConstructer struct {
	Activities []Activity
}

func NewActivity() *ActivityConstructer {
	return &ActivityConstructer{
		Activities: []Activity{},
	}
}

func (r *ActivityConstructer) AddActivity(activity Activity) {
	r.Activities = append(r.Activities, activity)
}

func (r *ActivityConstructer) GetActivities() []Activity {
	return r.Activities
}
