from __future__ import print_function
from ctypes import *

lib = cdll.LoadLibrary("T:\kana-reward\get_reward.so")

# define class GoSlice to map to:
# C type struct { void *data; GoInt len; GoInt cap; }
class GoSlice(Structure):
    _fields_ = [("data", POINTER(c_void_p)), ("len", c_longlong), ("cap", c_longlong)]

sparse = GoSlice((c_void_p * 29)(4, 6, 9, 11, 15, 24, 208, 286, 291, 308, 309, 328, 332, 345, 350, 353, 362, 366, 377, 381, 389, 394, 395, 409, 413, 442, 445, 449, 515), 29, 29) 
numeric = GoSlice((c_void_p * 6)(0, 0, 24000, 29000, 23000, 24000), 6, 6) 
progression = GoSlice((c_void_p * 71)(0, 157, 345, 581, 131, 281, 425, 569, 131, 297, 421, 593, 149, 277, 373, 583, 145, 265, 375, 523, 53, 173, 345, 565, 61, 241, 347, 519, 41, 253, 409, 952, 493, 19, 217, 391, 515, 125, 167, 335, 499, 15, 1318, 465, 145, 209, 319, 455, 135, 281, 417, 511, 147, 269, 423, 519, 113, 773, 293, 447, 535, 101, 237, 357, 587, 141, 207, 435, 527, 91, 289), 71, 71) 
option = GoSlice((c_void_p * 15)(0, 12, 16, 28, 32, 40, 48, 64, 68, 100, 104, 105, 106, 107, 108), 15, 15) 
action = GoSlice((c_void_p * 1)(12), 1, 1) 



# call Sort
lib.getReward.argtypes = [GoSlice, GoSlice, GoSlice, GoSlice, GoSlice]
lib.getReward.restype = c_longlong
lib.getReward(sparse, numeric, progression, option, action)
print("getReward = %d" % lib.getReward(sparse, numeric, progression, option, action))
