import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import '../localizations.dart';

class BottomNav extends StatelessWidget {
  const BottomNav({Key key, this.index}) : super(key: key);
  final int index;
  @override
  Widget build(BuildContext context) {
    final AppLocalizations_Labels labels = AppLocalizations.of(context);
    return BottomNavigationBar(
      items: <BottomNavigationBarItem>[
        buildWidget(
          icon: Icons.home,
          title: labels.home.news,
          context: context,
        ),
        buildWidget(
          icon: Icons.chat_bubble,
          title: labels.home.chat,
          context: context,
        ),
        buildWidget(
          icon: Icons.event_note,
          title: labels.home.enrollments,
          context: context,
        ),
      ],
      currentIndex: index,
      onTap: (int value) {},
    );
  }

  BottomNavigationBarItem buildWidget(
      {BuildContext context, IconData icon, String title}) {
    return BottomNavigationBarItem(
      activeIcon: Card(
        shape: const CircleBorder(),
        child: CircleAvatar(
          backgroundColor: Theme.of(context).cardColor,
          child: Icon(
            icon,
            color: Theme.of(context).colorScheme.secondary,
          ),
        ),
      ),
      icon: Icon(
        icon,
        color: Theme.of(context).primaryColorLight,
      ),
      title: Text(
        title,
        style: kIsWeb ||
                debugDefaultTargetPlatformOverride == TargetPlatform.fuchsia
            ? Theme.of(context).textTheme.subtitle
            : Theme.of(context)
                .textTheme
                .subtitle
                .copyWith(color: Theme.of(context).colorScheme.onPrimary),
      ),
    );
  }
}
