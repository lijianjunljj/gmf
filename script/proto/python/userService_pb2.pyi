import userModels_pb2 as _userModels_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UserDetailResponse(_message.Message):
    __slots__ = ["Code", "UserDetail"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    Code: int
    USERDETAIL_FIELD_NUMBER: _ClassVar[int]
    UserDetail: _userModels_pb2.UserModel
    def __init__(self, UserDetail: _Optional[_Union[_userModels_pb2.UserModel, _Mapping]] = ..., Code: _Optional[int] = ...) -> None: ...

class UserRequest(_message.Message):
    __slots__ = ["Password", "PasswordConfirm", "UserName"]
    PASSWORDCONFIRM_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    Password: str
    PasswordConfirm: str
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    UserName: str
    def __init__(self, UserName: _Optional[str] = ..., Password: _Optional[str] = ..., PasswordConfirm: _Optional[str] = ...) -> None: ...
