package mgogen

import (
	"gopkg.in/mgo.v2"
)

type Q struct {
	query      interface{}
	projection interface{}
	sort       []string
	skip       *int
	limit      *int
}

func (q *Q) Find(query interface{}) *Q {
	q.query = query
	return q
}

func (q *Q) Select(projection interface{}) *Q {
	q.projection = projection
	return q
}

func (q *Q) Sort(fields ...string) *Q {
	q.sort = fields
	return q
}

func (q *Q) Skip(skip int) *Q {
	if q.skip == nil {
		q.skip = new(int)
	}
	*q.skip = skip
	return q
}

func (q *Q) Limit(limit int) *Q {
	if q.limit == nil {
		q.limit = new(int)
	}
	*q.limit = limit
	return q
}

func (q *Q) ToQuery(session *mgo.Session, db, collection string) *mgo.Query {
	query := session.DB(db).C(collection).Find(q.query)
	query = query.Select(q.projection)
	if len(q.sort) > 0 {
		query = query.Sort(q.sort...)
	}
	if q.skip != nil {
		query = query.Skip(*q.skip)
	}
	if q.limit != nil {
		query = query.Limit(*q.limit)
	}
	return query
}
