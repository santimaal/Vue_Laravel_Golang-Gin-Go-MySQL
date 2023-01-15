<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\TableController;
use App\Http\Controllers\ThematicController;
use App\Http\Controllers\ReserveController;
use App\Http\Controllers\AuthController;

//All Users
// Thematic
Route::get('/thematic', [ThematicController::class, 'index']);
Route::get('/thematic/{id}', [ThematicController::class, 'show']);

// Table

Route::get('/table/{id}', [TableController::class, 'show']);

// User
Route::post('user/register', [AuthController::class, 'register']);
Route::post('user/login', [AuthController::class, 'login']);
Route::get('user/profile', [AuthController::class, 'GetProfile']);

// Unique Users ADMIN
Route::group(['middleware' => ['admin']], function () {

    // Table
    Route::get('/table', [TableController::class, 'index']);
    Route::post('/table', [TableController::class, 'store']);
    Route::put('/table/{id}', [TableController::class, 'update']);
    Route::delete('/table/{id}', [TableController::class, 'destroy']);
    // Route::resource('table', TableController::class);

    // Thematic
    Route::post('/thematic', [ThematicController::class, 'store']);
    Route::put('/thematic/{id}', [ThematicController::class, 'update']);
    Route::delete('/thematic/{id}', [ThematicController::class, 'destroy']);
    // Route::resource('thematic', ThematicController::class);

    // Reserve
    Route::get('/reserve', [ReserveController::class, 'index']);
    Route::get('/dashboard/reserves', [ReserveController::class, 'getReservesOrByP']);
    Route::put('/reserve/{id}', [ReserveController::class, 'update']);
    // Route::resource('reserve', ReserveController::class);

    //Users
    Route::resource('user', AuthController::class);
});
