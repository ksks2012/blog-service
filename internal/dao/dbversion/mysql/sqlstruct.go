// Code generated by go-literal-code-gen. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"

	metastore "github.com/semeqetjsakatayza/go-metastore-mysql"
)

const sqlCreateBlogTag = "CREATE TABLE `blog_tag` (" +
	"`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
	"`name` varchar(100) DEFAULT '' COMMENT '標籤名稱'," +
	"`created_on` int(10) unsigned DEFAULT '0' COMMENT '創建時間'," +
	"`created_by` varchar(100) DEFAULT '' COMMENT '創建人'," +
	"`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間'," +
	"`modified_by` varchar(100) DEFAULT '' COMMENT '修改人'," +
	"`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間'," +
	"`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除'," +
	"`state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用'," +
	"PRIMARY KEY (`id`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='標籤管理'"

const sqlCreateBlogArticleTag = "CREATE TABLE `blog_article_tag` (" +
	"`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
	"`article_id` int(11) NOT NULL COMMENT '文章ID'," +
	"`tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '標籤ID'," +
	"`created_on` int(10) unsigned DEFAULT '0' COMMENT '創建時間'," +
	"`created_by` varchar(100) DEFAULT '' COMMENT '創建人'," +
	"`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間'," +
	"`modified_by` varchar(100) DEFAULT '' COMMENT '修改人'," +
	"`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間'," +
	"`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除'," +
	"PRIMARY KEY (`id`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章標籤關聯'"

const sqlCreateBlogArticle = "CREATE TABLE `blog_article` (" +
	"`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
	"`title` varchar(100) DEFAULT '' COMMENT '文章標題'," +
	"`desc` varchar(255) DEFAULT '' COMMENT '文章簡述'," +
	"`cover_image_url` varchar(255) DEFAULT '' COMMENT '封面圖片地址'," +
	"`content` longtext COMMENT '文章内容'," +
	"`created_on` int(10) unsigned DEFAULT '0' COMMENT '新建時間'," +
	"`created_by` varchar(100) DEFAULT '' COMMENT '創建人'," +
	"`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間'," +
	"`modified_by` varchar(100) DEFAULT '' COMMENT '修改人'," +
	"`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間'," +
	"`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除'," +
	"`state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用'," +
	"PRIMARY KEY (`id`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理'"

const sqlCreateBlogAuth = "CREATE TABLE `blog_auth` (" +
	"`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
	"`app_key` varchar(20) DEFAULT '' COMMENT 'Key'," +
	"`app_secret` varchar(50) DEFAULT '' COMMENT 'Secret'," +
	"`created_on` int(10) unsigned DEFAULT '0' COMMENT '新建時間'," +
	"`created_by` varchar(100) DEFAULT '' COMMENT '創建人'," +
	"`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間'," +
	"`modified_by` varchar(100) DEFAULT '' COMMENT '修改人'," +
	"`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間'," +
	"`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除'," +
	"PRIMARY KEY (`id`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='認證管理'"

// ** SQL schema external filter

const metaKeyBlogServiceMetaSchemaRev = "blog-service-meta.schema"
const metaKeyBlogTagSchemaRev = "blog-tag.schema"
const metaKeyBlogArticleTagSchemaRev = "blog-article-tag.schema"
const metaKeyBlogArticleSchemaRev = "blog-article.schema"
const metaKeyBlogAuthSchemaRev = "blog-auth.schema"

const currentBlogServiceMetaSchemaRev = 2
const currentBlogTagSchemaRev = 1
const currentBlogArticleTagSchemaRev = 1
const currentBlogArticleSchemaRev = 1
const currentBlogAuthSchemaRev = 1

type schemaRevision struct {
	// TODO: unknown translation mode 0 for symbol [BlogServiceMeta]
	BlogTag        int32
	BlogArticleTag int32
	BlogArticle    int32
	BlogAuth       int32
}

func (rev *schemaRevision) IsUpToDate() bool {
	if currentBlogTagSchemaRev != rev.BlogTag {
		return false
	}
	if currentBlogArticleTagSchemaRev != rev.BlogArticleTag {
		return false
	}
	if currentBlogArticleSchemaRev != rev.BlogArticle {
		return false
	}
	if currentBlogAuthSchemaRev != rev.BlogAuth {
		return false
	}
	return true
}

type schemaManager struct {
	referenceTableName string
	ctx                context.Context
	conn               *sql.DB
}

func (m *schemaManager) FetchSchemaRevision() (schemaRev *schemaRevision, err error) {
	metaStoreInst := metastore.MetaStore{
		TableName: metaStoreTableName,
		Ctx:       m.ctx,
		Conn:      m.conn,
	}
	schemaRev = &schemaRevision{}
	if schemaRev.BlogTag, _, err = metaStoreInst.FetchRevision(metaKeyBlogTagSchemaRev); nil != err {
		return nil, err
	}
	if schemaRev.BlogArticleTag, _, err = metaStoreInst.FetchRevision(metaKeyBlogArticleTagSchemaRev); nil != err {
		return nil, err
	}
	if schemaRev.BlogArticle, _, err = metaStoreInst.FetchRevision(metaKeyBlogArticleSchemaRev); nil != err {
		return nil, err
	}
	if schemaRev.BlogAuth, _, err = metaStoreInst.FetchRevision(metaKeyBlogAuthSchemaRev); nil != err {
		return nil, err
	}
	return schemaRev, nil
}

func (m *schemaManager) updateBaseTableSchemaRevision(key string, rev int32) (err error) {
	metaStoreInst := metastore.MetaStore{
		TableName: metaStoreTableName,
		Ctx:       m.ctx,
		Conn:      m.conn,
	}
	err = metaStoreInst.StoreRevision(key, rev)
	return
}

func (m *schemaManager) execBaseSchemaModification(sqlStmt, schemaMetaKey string, targetRev int32) (err error) {
	if _, err = m.conn.Exec(sqlStmt); nil != err {
		return
	}
	return m.updateBaseTableSchemaRevision(schemaMetaKey, targetRev)
}

// upgrade routine for symbol not generated: BlogServiceMeta
func (m *schemaManager) UpgradeSchemaBlogTag(currentRev int32) (schemaChanged bool, err error) {
	switch currentRev {
	case currentBlogTagSchemaRev:
		return false, nil
	case 0:
		if err = m.execBaseSchemaModification(sqlCreateBlogTag, metaKeyBlogTagSchemaRev, currentBlogTagSchemaRev); nil == err {
			return true, nil
		}
	default:
		err = fmt.Errorf("unknown blog-tag schema revision: %d", currentRev)
	}
	return
}

func (m *schemaManager) UpgradeSchemaBlogArticleTag(currentRev int32) (schemaChanged bool, err error) {
	switch currentRev {
	case currentBlogArticleTagSchemaRev:
		return false, nil
	case 0:
		if err = m.execBaseSchemaModification(sqlCreateBlogArticleTag, metaKeyBlogArticleTagSchemaRev, currentBlogArticleTagSchemaRev); nil == err {
			return true, nil
		}
	default:
		err = fmt.Errorf("unknown blog-article-tag schema revision: %d", currentRev)
	}
	return
}

func (m *schemaManager) UpgradeSchemaBlogArticle(currentRev int32) (schemaChanged bool, err error) {
	switch currentRev {
	case currentBlogArticleSchemaRev:
		return false, nil
	case 0:
		if err = m.execBaseSchemaModification(sqlCreateBlogArticle, metaKeyBlogArticleSchemaRev, currentBlogArticleSchemaRev); nil == err {
			return true, nil
		}
	default:
		err = fmt.Errorf("unknown blog-article schema revision: %d", currentRev)
	}
	return
}

func (m *schemaManager) UpgradeSchemaBlogAuth(currentRev int32) (schemaChanged bool, err error) {
	switch currentRev {
	case currentBlogAuthSchemaRev:
		return false, nil
	case 0:
		if err = m.execBaseSchemaModification(sqlCreateBlogAuth, metaKeyBlogAuthSchemaRev, currentBlogAuthSchemaRev); nil == err {
			return true, nil
		}
	default:
		err = fmt.Errorf("unknown blog-auth schema revision: %d", currentRev)
	}
	return
}

// ** Generated code for 5 table entries
