package TokenService

import (
	MJ "../middleware"
	pb "../pb"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)


type TokenService struct {
	pb.UnimplementedTokenServiceServer
}

func (s *TokenService) mustEmbedUnimplementedTokenServiceServer() {
	panic("implement me")
}

func (s *TokenService) CreateToken(ctx context.Context,
	UserClaim *pb.UserClaim) (*pb.Token, error){
	fmt.Println("CreateToken service")
	//从massage里获取数据
	var claim = MJ.MyClaims{
		Username: UserClaim.Name,
		StandardClaims : jwt.StandardClaims{
			Audience: UserClaim.Audience,
			ExpiresAt: UserClaim.ExpiresAt,
			Id : UserClaim.Id,
			IssuedAt: UserClaim.IssuedAt,
			Issuer: UserClaim.Issuer,
			NotBefore: UserClaim.NotBefore,
			Subject: UserClaim.Subject,
		},
	}
	//编码，编码过程调用JWT-GO
	j := MJ.NewJWT()
	token, err := j.CreateToken(claim)

	//拼接返回数据
	TokenRes := pb.Token{Token: token}
	return &TokenRes, err
}

func (s *TokenService) ParserToken(ctx context.Context, Token *pb.Token) (*pb.UserClaim, error) {
	fmt.Println("ParserToken service")
	var token = Token.Token
	j := MJ.NewJWT()
	cll, err := j.ParserToken(token)
	var response pb.UserClaim
	if err != nil{
		response = pb.UserClaim{}
	}else {
		response = pb.UserClaim{
			Name: cll.Username,
			Audience  : cll.Audience,
			ExpiresAt: cll.StandardClaims.ExpiresAt,
			Id    : cll.Id,
			IssuedAt  :cll.IssuedAt,
			Issuer    : cll.Issuer,
			NotBefore : cll.NotBefore,
			Subject : cll.Subject,
		}
	}
	return &response, err
}

