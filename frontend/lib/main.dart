import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'screens/login_screen.dart';
import 'screens/main_screen.dart';

void main() {
  runApp(const ProviderScope(child: HiSeasApp()));
}

class HiSeasApp extends StatelessWidget {
  const HiSeasApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'HiSeas',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: const LoginScreen(),
      routes: {
        '/main': (context) => const MainScreen(),
      },
    );
  }
}
