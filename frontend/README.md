# HiSeas Flutter Frontend

This is the mobile client for the HiSeas project, built with Flutter.

## Features
- Riverpod for state management
- flutter_map for map display
- Basic login screen
- Main screen and map view

## Setup Instructions
1. Install [Flutter SDK](https://docs.flutter.dev/get-started/install)
2. In the `/frontend` directory, install dependencies:
   ```bash
   flutter pub get
   ```
3. Run the app:
   ```bash
   flutter run
   ```
   You can target Android, iOS, or web as supported by Flutter.

## Project Structure
- `lib/main.dart`: App entry point
- `lib/screens/login_screen.dart`: Login screen for API token
- `lib/screens/main_screen.dart`: Main screen
- `lib/screens/map_screen.dart`: Map view using flutter_map
- `lib/providers/auth_provider.dart`: Riverpod provider for authentication token

## Notes
- Update API endpoints and authentication logic in `lib/providers/auth_provider.dart` as needed to match your backend.
- The login screen expects an API token for authentication with the backend.
