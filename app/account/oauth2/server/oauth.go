package server

import (
	"chick/app/account/oauth2/model"

	"github.com/jinzhu/gorm"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func InitClient(db *gorm.DB) (clients []*model.OauthClient) {
	clients = make([]*model.OauthClient, 0)
	_ = db.Table("oauth_client").Select([]string{"key", "secret", "redirect_uri"}).Where("is_deleted = 0").Find(&clients).Error
	return
}

func NewOauthServer( /*redisCli *redis.ClusterClient, */ clients []*model.OauthClient) (oauthSrv *server.Server) {
	mgr := manage.NewDefaultManager()

	//mgr.MapTokenStorage(xStore.NewRedisClusterStoreWithCli(redisCli))
	mgr.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()

	for _, v := range clients {
		clientStore.Set(v.Key, &models.Client{
			ID:     v.Key,
			Secret: v.Secret,
			Domain: v.RedirectUri,
		})
	}

	mgr.MapClientStorage(clientStore)

	mgr.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	return server.NewDefaultServer(mgr)
}
