import os
import json
import time

from config import DATA_PATH


class Passwords():
    def __init__(self):
        if not os.path.exists(DATA_PATH):
            os.makedirs(DATA_PATH)
        if not os.path.exists(f'{DATA_PATH}/passwords.json'):
            with open(f'{DATA_PATH}/passwords.json', 'w') as f:
                f.write('{}')
            self.passwords = {}
            return
        self.password_verifications = {}
        with open(f'{DATA_PATH}/passwords.json', 'r') as f:
            self.passwords = json.load(f)

    def get(self, user):
        return self.passwords.get(user)

    def set(self, user, password):
        verification_code = os.urandom(3).hex()
        self.password_verifications[verification_code] = {
            'user': user,
            'password': password,
            'code_expire': 300 + int(time.time())
        }
        return verification_code

    def verify(self, user, verification_code):
        if verification_code not in self.password_verifications:
            return False, 'Verification code not found'
        if self.password_verifications[verification_code]['user'] != user:
            return False, 'Verification code does not match user'
        if self.password_verifications[verification_code]['code_expire'] < int(time.time()):
            return False, 'Verification code expired'
        self.passwords[user] = {
            'password': self.password_verifications[verification_code]['password'],
            'active': True,
        }
        del self.password_verifications[verification_code]
        with open(f'{DATA_PATH}/passwords.json', 'w') as f:
            json.dump(self.passwords, f)
        return True, 'Password set successfully'
