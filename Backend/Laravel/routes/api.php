<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\TableController;
use App\Http\Controllers\ThematicController;
use App\Http\Controllers\UserController;
use App\Http\Controllers\ReserveController;
use App\Http\Controllers\AuthController;


// Table
Route::get('/table', [TableController::class, 'index']);
Route::get('/table/{id}', [TableController::class, 'show']);
Route::post('/table', [TableController::class, 'store']);
Route::put('/table/{id}', [TableController::class, 'update']);
Route::delete('/table/{id}', [TableController::class, 'destroy']);
// Route::resource('table', TableController::class);

// Thematic
Route::get('/thematic', [ThematicController::class, 'index']);
Route::get('/thematic/{id}', [ThematicController::class, 'show']);
Route::post('/thematic', [ThematicController::class, 'store']);
Route::put('/thematic/{id}', [ThematicController::class, 'update']);
Route::delete('/thematic/{id}', [ThematicController::class, 'destroy']);
// Route::resource('thematic', ThematicController::class);


// User
Route::post('user/register', [AuthController::class, 'register']);
Route::post('user/login', [AuthController::class, 'login']);
Route::resource('user', AuthController::class);



// Reserves
Route::resource('reserve', ReserveController::class);


