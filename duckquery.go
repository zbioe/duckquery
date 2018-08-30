package duckquery

type Output struct {
	Abstract	string
	Labels	[]Label
	Topics	[]string
}

type Label struct {
	Label	string
	Value	string
}

func Query(s string) (Output, error) {
	var o Output
	ao, err := Request(s)
	if err != nil {
		return o, err
	}
	o.Abstract = ao.Abstract
	for _, t := range ao.RelatedTopics {
		o.Topics = append(o.Topics, t.Text)
	}
	for _, l := range ao.Infobox.Content {
		if l.Value != "" && l.Label != "" {
			o.Labels = append(o.Labels, l)
		}
	}
	return o, nil
}
