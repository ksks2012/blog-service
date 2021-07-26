package mysql

import (
	bsstorages "github.com/blog-service/internal/dao/dbversion"
)

func (m *schemaManager) UpgradeSchema(currentRevs *schemaRevision) (schemaChanged bool, err error) {
	status := bsstorages.SchemaUpgradeStatus{
		Changed:   false,
		LastError: nil,
	}
	status.RunUpgrade("blog-tag", m.UpgradeSchemaBlogTag, currentRevs.BlogTag)
	status.RunUpgrade("blog-article-tag", m.UpgradeSchemaBlogArticleTag, currentRevs.BlogArticleTag)
	status.RunUpgrade("blog-article", m.UpgradeSchemaBlogArticle, currentRevs.BlogArticle)
	return status.Changed, status.LastError
}
