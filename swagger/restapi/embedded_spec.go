// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Splathonで使うAPI",
    "title": "SplathonAPI",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/splathon/",
  "paths": {
    "/v{eventId}/matches/{matchId}": {
      "get": {
        "description": "マッチの詳細を返す。スコアボードとかで使える。",
        "tags": [
          "match"
        ],
        "operationId": "getMatch",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "eventId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "match id",
            "name": "matchId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Match"
            }
          },
          "default": {
            "description": "Generic error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v{eventId}/results": {
      "get": {
        "description": "リザルト一覧を返す。リザルトと言いつつ終了していない未来のマッチも返す。ゲスト・管理アプリ両方から使う。team_idを指定するとそのチームのみの結果が返ってくる。",
        "tags": [
          "result"
        ],
        "operationId": "getResult",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "eventId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "team id",
            "name": "team_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Results"
            }
          },
          "default": {
            "description": "Generic error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Battle": {
      "description": "バトル。勝敗などは決まってない状態のこともある。",
      "type": "object",
      "required": [
        "id",
        "order"
      ],
      "properties": {
        "id": {
          "description": "Battle ID",
          "type": "integer",
          "format": "int32"
        },
        "order": {
          "description": "何戦目か",
          "type": "integer",
          "format": "int32"
        },
        "rule": {
          "type": "object",
          "properties": {
            "key": {
              "description": "Rule key. ref: https://splatoon2.ink/data/locale/ja.json",
              "type": "string",
              "enum": [
                "turf_war",
                "splat_zones",
                "tower_control",
                "rainmaker",
                "clam_blitz"
              ]
            },
            "name": {
              "description": "Localized rule name.",
              "type": "string"
            }
          }
        },
        "stage": {
          "type": "object",
          "properties": {
            "id": {
              "description": "Stage ID. ref: https://splatoon2.ink/data/locale/ja.json",
              "type": "integer",
              "format": "int32"
            },
            "name": {
              "description": "Localized stage name.",
              "type": "string"
            }
          }
        },
        "winner": {
          "description": "勝者がどちらか。",
          "type": "string",
          "enum": [
            "alpha",
            "bravo"
          ]
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Match": {
      "type": "object",
      "required": [
        "id",
        "teamAlpha",
        "teamBravo"
      ],
      "properties": {
        "battles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Battle"
          }
        },
        "id": {
          "description": "Match ID",
          "type": "integer",
          "format": "int32"
        },
        "order": {
          "description": "Room内でのマッチの順番",
          "type": "integer",
          "format": "int32"
        },
        "teamAlpha": {
          "$ref": "#/definitions/Team"
        },
        "teamBravo": {
          "$ref": "#/definitions/Team"
        },
        "winner": {
          "description": "勝者がどちらか。または引き分け。",
          "type": "string",
          "enum": [
            "alpha",
            "bravo",
            "draw"
          ]
        }
      }
    },
    "Member": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "icon": {
          "description": "Slack icon URL",
          "type": "string"
        },
        "id": {
          "description": "Member ID (Slack ID かも?)",
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Results": {
      "description": "予選/決勝Tのリザルト。予選/決勝Tは同じ構造なのでフラットにできるがクライアントがトーナメント表だせる拡張性持たせるために別フィールドで持つ。",
      "type": "object",
      "properties": {
        "qualifiers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Round"
          }
        },
        "tournament": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Round"
          }
        }
      }
    },
    "Room": {
      "description": "Roomごとのマッチ",
      "type": "object",
      "required": [
        "name",
        "matches"
      ],
      "properties": {
        "id": {
          "description": "Room ID.",
          "type": "integer",
          "format": "int32"
        },
        "matches": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Match"
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Round": {
      "description": "予選/決勝Tラウンド両方扱う。",
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "description": "ラウンド名。e.g. 予選第1ラウンド, 決勝T1回戦, 決勝戦",
          "type": "string"
        },
        "rooms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Room"
          }
        },
        "round": {
          "description": "何ラウンドか。i.e. 予選第Nラウンド, 決勝T N回戦",
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "Team": {
      "type": "object",
      "required": [
        "id",
        "name",
        "companyName"
      ],
      "properties": {
        "companyName": {
          "type": "string"
        },
        "id": {
          "description": "Team ID",
          "type": "integer",
          "format": "int32"
        },
        "members": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Member"
          }
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "リザルト一覧",
      "name": "result"
    },
    {
      "description": "マッチ",
      "name": "match"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Splathonで使うAPI",
    "title": "SplathonAPI",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/splathon/",
  "paths": {
    "/v{eventId}/matches/{matchId}": {
      "get": {
        "description": "マッチの詳細を返す。スコアボードとかで使える。",
        "tags": [
          "match"
        ],
        "operationId": "getMatch",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "eventId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "match id",
            "name": "matchId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Match"
            }
          },
          "default": {
            "description": "Generic error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v{eventId}/results": {
      "get": {
        "description": "リザルト一覧を返す。リザルトと言いつつ終了していない未来のマッチも返す。ゲスト・管理アプリ両方から使う。team_idを指定するとそのチームのみの結果が返ってくる。",
        "tags": [
          "result"
        ],
        "operationId": "getResult",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "eventId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "team id",
            "name": "team_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Results"
            }
          },
          "default": {
            "description": "Generic error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Battle": {
      "description": "バトル。勝敗などは決まってない状態のこともある。",
      "type": "object",
      "required": [
        "id",
        "order"
      ],
      "properties": {
        "id": {
          "description": "Battle ID",
          "type": "integer",
          "format": "int32"
        },
        "order": {
          "description": "何戦目か",
          "type": "integer",
          "format": "int32"
        },
        "rule": {
          "type": "object",
          "properties": {
            "key": {
              "description": "Rule key. ref: https://splatoon2.ink/data/locale/ja.json",
              "type": "string",
              "enum": [
                "turf_war",
                "splat_zones",
                "tower_control",
                "rainmaker",
                "clam_blitz"
              ]
            },
            "name": {
              "description": "Localized rule name.",
              "type": "string"
            }
          }
        },
        "stage": {
          "type": "object",
          "properties": {
            "id": {
              "description": "Stage ID. ref: https://splatoon2.ink/data/locale/ja.json",
              "type": "integer",
              "format": "int32"
            },
            "name": {
              "description": "Localized stage name.",
              "type": "string"
            }
          }
        },
        "winner": {
          "description": "勝者がどちらか。",
          "type": "string",
          "enum": [
            "alpha",
            "bravo"
          ]
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Match": {
      "type": "object",
      "required": [
        "id",
        "teamAlpha",
        "teamBravo"
      ],
      "properties": {
        "battles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Battle"
          }
        },
        "id": {
          "description": "Match ID",
          "type": "integer",
          "format": "int32"
        },
        "order": {
          "description": "Room内でのマッチの順番",
          "type": "integer",
          "format": "int32"
        },
        "teamAlpha": {
          "$ref": "#/definitions/Team"
        },
        "teamBravo": {
          "$ref": "#/definitions/Team"
        },
        "winner": {
          "description": "勝者がどちらか。または引き分け。",
          "type": "string",
          "enum": [
            "alpha",
            "bravo",
            "draw"
          ]
        }
      }
    },
    "Member": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "icon": {
          "description": "Slack icon URL",
          "type": "string"
        },
        "id": {
          "description": "Member ID (Slack ID かも?)",
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Results": {
      "description": "予選/決勝Tのリザルト。予選/決勝Tは同じ構造なのでフラットにできるがクライアントがトーナメント表だせる拡張性持たせるために別フィールドで持つ。",
      "type": "object",
      "properties": {
        "qualifiers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Round"
          }
        },
        "tournament": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Round"
          }
        }
      }
    },
    "Room": {
      "description": "Roomごとのマッチ",
      "type": "object",
      "required": [
        "name",
        "matches"
      ],
      "properties": {
        "id": {
          "description": "Room ID.",
          "type": "integer",
          "format": "int32"
        },
        "matches": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Match"
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Round": {
      "description": "予選/決勝Tラウンド両方扱う。",
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "description": "ラウンド名。e.g. 予選第1ラウンド, 決勝T1回戦, 決勝戦",
          "type": "string"
        },
        "rooms": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Room"
          }
        },
        "round": {
          "description": "何ラウンドか。i.e. 予選第Nラウンド, 決勝T N回戦",
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "Team": {
      "type": "object",
      "required": [
        "id",
        "name",
        "companyName"
      ],
      "properties": {
        "companyName": {
          "type": "string"
        },
        "id": {
          "description": "Team ID",
          "type": "integer",
          "format": "int32"
        },
        "members": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Member"
          }
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "リザルト一覧",
      "name": "result"
    },
    {
      "description": "マッチ",
      "name": "match"
    }
  ]
}`))
}
