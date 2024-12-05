import json
import urllib.request
import firebase_admin
from firebase_admin import credentials, auth
import os

def lambda_handler(event, context):
    # 環境変数からサービスアカウントキーJSONを読み込む
    service_account_key_str = os.environ.get('SERVICE_ACCOUNT_KEY')
    if not service_account_key_str:
        raise ValueError('Environment variable SERVICE_ACCOUNT_KEY is not set')

    # JSON文字列を辞書型に変換
    service_account_key = json.loads(service_account_key_str)

    # Firebase Admin SDKの初期化
    cred = credentials.Certificate(service_account_key)
    try:
        firebase_admin.initialize_app(cred)
    except ValueError:
        pass  # すでに初期化されている場合

    # 環境変数から特定のユーザーのUIDを取得
    uid = os.environ.get('USER_UID')
    if not uid:
        raise ValueError('Environment variable USER_UID is not set')

    # カスタムトークンの生成
    custom_token = auth.create_custom_token(uid)

    # カスタムトークンをIDトークンに交換
    id_token = get_id_token(custom_token)

    # 環境変数からAPIのURLを取得
    api_url = os.environ.get('API_URL')
    if not api_url:
        raise ValueError('Environment variable API_URL is not set')

    # APIを叩く
    headers = {
        'Authorization': f'Bearer {id_token}',
        'Content-Type': 'application/json'
    }
    request = urllib.request.Request(api_url, headers=headers, method='GET')

    try:
        with urllib.request.urlopen(request) as response:
            result = response.read()
            print('API Response:', result)
    except Exception as e:
        print('Error calling API:', e)
        raise e

    return {
        'statusCode': 200,
        'body': json.dumps('Function executed successfully!')
    }

def get_id_token(custom_token):
    # 環境変数からFirebase APIキーを取得
    firebase_api_key = os.environ.get('FIREBASE_API_KEY')
    if not firebase_api_key:
        raise ValueError('Environment variable FIREBASE_API_KEY is not set')

    # カスタムトークンをIDトークンに交換するエンドポイント
    url = f"https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key={firebase_api_key}"
    data = json.dumps({
        'token': custom_token.decode('utf-8'),
        'returnSecureToken': True
    }).encode('utf-8')

    request = urllib.request.Request(url, data=data, headers={'Content-Type': 'application/json'})
    try:
        with urllib.request.urlopen(request) as response:
            response_data = json.loads(response.read())
            id_token = response_data['idToken']
            return id_token
    except Exception as e:
        print('Error exchanging custom token:', e)
        raise e
